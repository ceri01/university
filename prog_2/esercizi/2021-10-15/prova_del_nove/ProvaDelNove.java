import java.util.*;

public class ProvaDelNove {
    static int sumDigit(int n) {
        int t = 0;
        String m1 = String.valueOf(n); 
        
        for(int i = 0; i < m1.length(); i++) {
        t += m1.charAt(i) - '0'; 
            // t += Character.getNumericValue(m1.charAt(i)); per farlo con gli oggetti
        }

        return t;
    }

    static boolean proof(int mdo, int mre, int res) {
        int tot;
        tot = sumDigit(mdo) * sumDigit(mre);
        for(;tot > 9;) {
            tot = sumDigit(tot);
        }
        for(;res > 9;) {
            res = sumDigit(res);
        }
        System.out.print(tot);
        System.out.println(res);
        if (tot == res) {
            return true;
        }
        return false;
    }
    
    public static void main(String[] args) {
        Scanner s =  new Scanner(System.in);
        int n = s.nextInt();
        int j = 1;
        
        for(int i = 1; i < n + 1; i++) {
            for(int k = 1; k < n + 1; k++) {
                if (k * i != j) {
                    if (proof(i, k, j)) {
                        System.out.println(String.valueOf(i) + " " + String.valueOf(k) + " " + String.valueOf(j));
                    }
                    j++;
                }
            }
            j = 1;
        }
    }
}
