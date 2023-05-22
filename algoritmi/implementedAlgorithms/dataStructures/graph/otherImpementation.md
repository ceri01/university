# Altre implementazioni dei grafi

Esistono altre implementazioni dei grafi, come la lista d'incidenza e la matrice d'incidenza, ma sono state solo presentate
e mai effettivamente durante il corsi di Algoritmi e Strutture dati (2021/2022).
Di seguito è presente una breve spiegazione di queste due implementazioni.

## Lista d'incidenza

La lista d'incidenza consiste in una struttura dati formata dall'insieme degli archi (coppie di nodi) e da una lista in cui
per ogni elemento è presente una lista di puntatori agli archi in cui il nodo in questione conpare.
Questa rappresentazione ha un costo in termini di spazio nell'ordine di θ(n+m), infatti la dimensione della lista dei puntatori (m)
dipende dalla dimensione della lista degli archi (n).
Questa implementazione risente degli stessi problemi della lista di adiacenza, ovvero che per un grafo orientato è costoso
capire quali sono gli archi entranti in un nodo.

## Matrice d'incidenza