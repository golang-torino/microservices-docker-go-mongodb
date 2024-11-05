Ora che hai clonato il repo, che fare?

1. Assicurati tutto funzioni. Lancia `make start` e quando e' tutto pronto visita [localhost](http://localhost). Devi vedere un sito, "Cinema Website".
2. Ispeziona i vari tool di observability locali a disposizione, dovresti vedere tracce, metriche e log! Seguie [docs/observability](docs/observability.md).
3. Clicca in giro e ispeziona quello che viene raccolto e come viene visualizzato. Quando ti senti pront* passa ai task qui sotto.

## Tasklist

- Non tutti i servizi sono instrumentati! Obiettivo: instrumenta un servizio collezionando le tracce dal server
  - Le tracce del server da sole non bastano, aggiungi delle tracce specifiche per ispezionare la comunicazione con il database.
- Perche' nella traccia `website: /users/view/{id}` viene riportato due volte lo span `HTTP GET`? Cosa sta succedendo?
- Non raccogliamo alcuna metrica dal nostro DB (MongoDB). Fai felice il DBA che e' in te, aggiungi la collezione di metriche e logs.
- E' disponibile l'instrumentazione automatica con eBPF, perche' non provarla? https://github.com/open-telemetry/opentelemetry-go-instrumentation
