typedef struct graph *Graph;

/* crea un grafo con n nodi */
Graph graph_new( int n );

/* distrugge il grafo g */
void graph_destroy( Graph g );

/* inserisce l'arco (v,w) nel grafo g */
void graph_edge_insert( Graph g, int v, int w );

/* legge da standard input un grafo
 * il grafo viene creato e restituito
 * il primo intero inserito indica il numero di nodi, il secondo invece indica il numero di archi che l'utente vuole inserire.
 * l'utente deve inserire per ogni riga i due nodi che saranno collegati da un arco.
 */
Graph graph_read(void);

/* stampa su standard output un grafo
 * vengoono stampati gli archi tra i due nodi
 */
void graph_print( Graph g );

// Permette di effettuare la visira in profondita' di un grafo

void dfs(Graph g);

// Funzione usata in dfs per effettuare la ricorsione, l'utente non la deve usare

void dfs_rec(Graph g, int i, int *vis);

// Permette di effettuare la visira in ampiezza di un grafo

void bfs(Graph g);
