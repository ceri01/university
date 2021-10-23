import java.util.*;

public class Kaprekar {
    
    // Pre-condizioni: a non null
    // Post-condizioni: Al termine della procedura conterrà il suo reverse
    // Effetti-collaterali: a potrebbe essere modificato
    static void reverse(byte[] a) {
        for(int i = 0; i < (a.length/2); i++) {
            byte tmp = a[i];
            a[i] = a[a.length - 1 - i];
            a[a.length - 1 - i] = tmp;
        }
    }  
    
    // Pre-condizioni: a non null
    // Post-condizioni: Al termine della procedura il contenuto dell'array verrà sortato
    // Effetti-collaterali: a potrebbe essere modificato
    static void sort(byte[] a) {
        // fai quello specificato dal prof
        // intanto usiamo quelle di java
        

    }

    // Pre-condizioni: a non null, ciascuna posizione di a deve contenere un numero compreso tra 0 e 9
    // Post-condizioni: Restituisce l'intero le cui cifre sono rappresentate dai numeri interi memorizzati da a
    static int digitsToNum(byte[] a) {
        int n = 0;
        for (byte b : a) {
            n = n * 10 + b;
        }
        return n;
    }

    // Pre-condizioni: 0 < digits, digits <= cifre di n, n >= 0
    // Post-condizioni:i restituisce un array contenente le cifre di n
    static byte[] numToDigits(int n, int digits) {
        byte[] a = new byte[digits];
        for(int i = 0; i < digits; i++) {
            a[i] = (byte) (n % 10);
            n /= 10;
        }
        return a;
    }

    // Pre-condizioni: n > 0 
    // Post-condizioni: restituire il numero successivo della sequenza di kaprekr
    static int nextKaprekar(int n) {
        byte[] a = numToDigits(n, 4);
        Arrays.sort(a);
        n = -digitsToNum(a);
        reverse(a);
        n += digitsToNum(a);
        return n;
    }

    // Pre-condizioni: 0 <= n <= 9999, n deve avere almeno due cifre diverse
    // Post-condizioni: Stampa la sequenza di kaprekar a partire dal nimero n
    // Effetti-collaterali: il flusso standard di output viene modificato
    static void printKaprekar(int n) {
        if (n < 0 || n > 9999) throw new IllegalArgumentException("L'input deve essere compreso tra 0 e 9999");

        while(n != 6174) {
            System.out.println(n);
            n = nextKaprekar(n);
        }
        System.out.println(n);
    }


    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        int n = s.nextInt();
        printKaprekar(n);
    }
}
