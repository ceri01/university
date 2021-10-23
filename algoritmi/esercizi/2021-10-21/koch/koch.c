#include "libpsgraph.h"

void disegna(int i, int x) {
    if (i == 0) {
        draw(x);
        return;    
    } else {
        disegna(i - 1, x / 3);
        turn(-60);   
        disegna(i - 1, x / 3);
        turn(+120);
        disegna(i - 1, x / 3);
        turn(-60);
        disegna(i - 1, x / 3);
    }
}

int main() {
    start("disegno.ps");
    disegna(2, 60);
    end();
}
