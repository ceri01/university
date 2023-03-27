import java.util.*;

public class Scacchiera {
    
    // Pre-condizioni: n > 0
    // Post-condizioni: Stampa una riga n*n  della scacchiera 
    static void generaRiga(int n, int ch) {
        for(int r = 0; r < n; r++) {
            for (int i = 0; i < 8; i++) {
                for (int j = 0; j < n; j++) {
                    System.out.print((char)ch);
                }
                if(ch == 45) {
                    ch = ch - 10;
                } else {
                    ch = ch + 10;
                }
            }
            System.out.println("");
        }   
    }
    

    // Pre-condizioni: n > 0
    // Post-condizioni: Stampa una scacchiera
    static void stampaScacchiera(int n) {
        int ch = 45;
        for(int i = 0; i < 8; i++) {
            generaRiga(n, ch);
            if(ch == 45) {
                ch -= 10;
            } else  {
                ch += 10;
            }
        }
    }      
    
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
        stampaScacchiera(n);
    }
}  
