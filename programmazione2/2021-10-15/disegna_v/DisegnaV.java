import java.util.*;

public class DisegnaV {
    
    // Pre-condizioni: n > 0
    // Post-condizioni: Stampa una v di dimensione n
    static void disegna(int n) {
        int spazi = n + (n - 3);
        for(int i = 0; i < n; i++) {
            for(int k = 0; k <= spazi + (i + 1);k++) {
                if(k == i || k == spazi + (i + 1)) {
                    System.out.print("*");
                } else {
                    System.out.print(" ");
                }
            }
            spazi -= 2;
            System.out.println("");
        }
    }
    
    public static void main(String[] args) {
        Scanner s =  new Scanner(System.in);
        int n = s.nextInt();
        disegna(n);
    }
}
