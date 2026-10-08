Creami uno spec.md per un software che
1. data un'immagine (vera o nanobanana) crea una versione 3d (tipo tripo3d) del sistema e un arnese per ruotarla.
2. volendo dato un prompt fa direttamente tutto: immagine -> coso 3d.
La cosa 3d dev'essere compatibile con blender, e poi godot e game engines, se umana deve potersi adattare a uno skeleton. Magari data una libreria di skeleton (eg umani) deve poter essere testata su quegli scheletri.

Il suistema deve avere un feedback loop sofisticato in modo che il modello 3d puo' essere testato o da solo (rigido) o su skeleton, in maniera che se ci sono errori il sistema lo puo editare deterministicamente o con EVAL. Insomma, l'idea e' che io ti do un'immagine, tu crei una versione 3d, e il sistema ha un modo di dire 'fa schifo per la ragione X Y Z ->' questo output puo' essere dato in pasto a una LLM che usa quell'info per re-iterare (chesso, i piedi sono piccoli, la faccia e' tutta sbagliata, il naso e' piatto, o visto da questa angolazione non somiglia per niente all'immagine...) e cosi via.

Il tutto avra' una UI e un editor che (dopo aver lavorato per qualche minuto) mi da la versione 3d auto-rotante. E ovviamente ha una galleria per salvare il tutto.

---

Follow-ups (same session):

1. pensavo di fare tutto in locale! Su pupurabbu che ha GPU e magari su cloud run una volta dimostrato che funge. voglio GRATIS e lavoro per google quindi cloud run + GPU e' gratis per me :) ma testiamolo prima in locale per il coding feedback loop
2. boh sono ignorante. un paio di skeleton free e umani.
3. 10min si, ZERO dollari. gratis.
4. non c'e', ci sara', faremo un palladius/image23d o image2model3d decidi tu :)
PS sono in svizzera non EU, se cambia.
