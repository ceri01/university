import java.util.*;

public class ProvaDelNove {
    
    // Pre-condizioni: n > 0
    // Post-condizioni: Ritorna la somma delle cifre di un numero
    static int sommaCifre(int n) {
        return n % 9;
    }
    
    // Pre-condizioni: mdo > 0, mre > 0, res > 0
    // Post-condizioni: ritorna true se mdo * mre e res soddisfano la prova del nove 
    static boolean prova(int mdo, int mre, int res) {
        int tot;
        tot = sommaCifre(mdo) * sommaCifre(mre);
        for(;tot > 9;) {
            tot = sommaCifre(tot);
        }
        for(;res > 9;) {
            res = sommaCifre(res);
        }
        if (tot == res) {
            return true;
        }
        return false;
    }
    
    // Pre-condizioni: n > 0;
    // Post-condizioni: Stampa tutte le terne tali che A * B != C (con A, B e C < n) e che soddisfano la prova del nove
    static void stampaTerne(int n) {
        for(int a = 1; a < n; a++) {
            for(int b = 1; b < n; b++) {
                for(int c = 1; c < n; c++) {
                    if (a * b != c && prova(a, b, c)) {
                        System.out.println(String.valueOf(a) + " " + String.valueOf(b) + " " + String.valueOf(c));
                    }
                }
            }
        }
    }
    
    public static void main(String[] args) {
        Scanner s =  new Scanner(System.in);
        int n = s.nextInt(); 
        stampaTerne(n);
    }
}
