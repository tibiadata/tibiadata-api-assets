package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
)

func testClient() *resty.Client {
	return resty.New()
}

func withTibiaServer(t *testing.T, handler http.Handler) {
	t.Helper()

	server := httptest.NewServer(handler)
	oldHost, oldProtocol := TibiaComHost, TibiaComProtocol
	TibiaComHost, TibiaComProtocol = server.Listener.Addr().String(), "http"
	t.Cleanup(func() {
		TibiaComHost, TibiaComProtocol = oldHost, oldProtocol
		server.Close()
	})
}

func TestCreaturesWorker(t *testing.T) {
	withTibiaServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/library/" || r.URL.Query().Get("subtopic") != "creatures" {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte(`<div class="BoxContent"><div class="Creatures">
			<div><div>Rats</div><a href="?subtopic=creatures&race=rat">Rats</a></div>
			<div><div>Ravenous Lava Lurkers</div><a href="?subtopic=creatures&race=aggressivelavalurker">Ravenous Lava Lurkers</a></div>
		</div></div>`))
	}))

	var builder Builder
	if err := builder.creaturesWorker(testClient()); err != nil {
		t.Fatalf("creaturesWorker() error = %v", err)
	}

	want := []Creature{
		{Endpoint: "rat", PluralName: "Rats", Name: "Rat"},
		{Endpoint: "aggressivelavalurker", PluralName: "Ravenous Lava Lurkers", Name: "Ravenous Lava Lurkers"},
	}
	if len(builder.Creatures) != len(want) {
		t.Fatalf("got %d creatures, want %d", len(builder.Creatures), len(want))
	}
	for i := range want {
		if builder.Creatures[i] != want[i] {
			t.Errorf("creature %d = %#v, want %#v", i, builder.Creatures[i], want[i])
		}
	}
}

func TestCreaturesWorkerRejectsUnexpectedURL(t *testing.T) {
	withTibiaServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<div class="BoxContent"><div class="Creatures">
			<div><div>Rats</div><a href="?subtopic=creatures">Rats</a></div>
		</div></div>`))
	}))

	var builder Builder
	if err := builder.creaturesWorker(testClient()); err == nil {
		t.Fatal("creaturesWorker() succeeded for a creature URL without a race")
	}
}

func TestSpellsWorker(t *testing.T) {
	withTibiaServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<div class="Table3"><table class="TableContent">
			<tr><th>Spell</th></tr>
			<tr><td>Energy Wave (exevo vis hur)</td></tr>
			<tr><td>Apprentice's Strike (exori min)</td></tr>
		</table></div>`))
	}))

	var builder Builder
	if err := builder.spellsWorker(testClient()); err != nil {
		t.Fatalf("spellsWorker() error = %v", err)
	}

	want := []Spell{
		{Name: "Energy Wave", Formula: "exevo vis hur", Endpoint: "energywave"},
		{Name: "Apprentice's Strike", Formula: "exori min", Endpoint: "apprenticestrike"},
	}
	if len(builder.Spells) != len(want) {
		t.Fatalf("got %d spells, want %d", len(builder.Spells), len(want))
	}
	for i := range want {
		if builder.Spells[i] != want[i] {
			t.Errorf("spell %d = %#v, want %#v", i, builder.Spells[i], want[i])
		}
	}
}

func TestSpellsWorkerRejectsMalformedSpell(t *testing.T) {
	withTibiaServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<div class="Table3"><table class="TableContent">
			<tr><th>Spell</th></tr><tr><td>Energy Wave</td></tr>
		</table></div>`))
	}))

	var builder Builder
	if err := builder.spellsWorker(testClient()); err == nil {
		t.Fatal("spellsWorker() succeeded for malformed spell text")
	}
}

func TestGetDocRejectsNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	if _, err := getDoc(testClient(), server.URL); err == nil {
		t.Fatal("getDoc() succeeded for a non-200 response")
	}
}
