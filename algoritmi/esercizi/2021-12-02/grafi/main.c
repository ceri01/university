#include "grafi.h"

int main() {
    Graph g = graph_read();
    dfs(g);
    //graph_print(g);
    graph_destroy(g);
}
