#include "grafi.h"

int main() {
    Graph g = graph_read();
    graph_print(g);
    graph_destroy(g);
}
