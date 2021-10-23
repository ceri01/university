import java.util.*;

public class Scacchiera {
    static void generaRiga(int n, char ch) {
        for (int i = 0; i < (n * 8); i++) {
            for (int j = 0; j < n; j++) {
                System.out.print(ch);
            }
        }
    }

    static void stampaScacchiera(int n) {
        int ch = 45;
        for (int i = 0; i < 8; i++) {
            generaRiga(n, (char)ch);
            if (i % 2 == 1) {
                ch += 10;
            } else {
                ch -= 10;
            }
        }
    }      
    
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
        stampaScacchiera(n);
    }
}  
