# Input Prompt (Verbatim Voice Notes)

## Riccardo's Vision
"brainstorming il mio sogno di fare valutazioni e loggare traiettorie sull'uso di gemini in gcp da mille client da mille posti da mille project di diverse api key e loggarli in un unico punto centralizzato e trasparente tipo in un log sync di cloud logging... se ne implementassimo un proxy trasparente che tu in dev lo usi su localhost e in prod su cloud run e che in entrambi i casi dà ampi metadati e crea tracce per ogni conversazione... in locale hai un frontend javascript che appena parte una traccia pum... ti dice token io, modello chiamato, latenza..."

"un mio collega a google ha fatto un proxy in go che intercettava tutte le sfiltrazioni di carte di credito e api keys client side... deve essere implementato in go o rust e non deve aumentare la latenza di più di 10-30 ms... 70 ms se devi fare decisioni ci sta, ma 1-2 secondi no..."

"un router che intercetta il traffico, poi lo manda di là e si greppa metadati molto piccoli... crea log e traces per analizzare il testo. Per audio/musica/multimodalità ci sarà un link a GCS o logga solo il testo di default per restare leggeri..."
