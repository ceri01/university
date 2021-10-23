import java.util.*;

public class Rombo {
    static void generaRombo(int n) {
        int diag = (2 * n) + 1;
        int p = 1;
        int tmpN = n;
        for (int i = 0; i < diag; i++) {
            for(int k = 0; k < p + tmpN; k++) {
                if (k < tmpN) {
                    System.out.print(" ");
                } else {
                    System.out.print("*");
                }
            }
            if(i < n) {
                p += 2;
                tmpN--;
            } else {
                p -= 2;
                tmpN++;
            }
            System.out.println("");
        }
    }
    
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
        generaRombo(n);
    }
}
