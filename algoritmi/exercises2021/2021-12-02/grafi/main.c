#include "grafi.h"

int main() {
    Graph g = graph_read();
    bfs(g);
    //graph_print(g);
    graph_destroy(g);
}
