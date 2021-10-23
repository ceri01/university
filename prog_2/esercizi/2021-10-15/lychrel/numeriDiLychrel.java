import java.util.*;

public class numeriDiLychrel {
    
    // Pre-condizioni: s not null
    // Post-condizioni: restituisce il reverse di s
    // Effetti collaterali: nessuno   
    static String reverse(String s) {
        int len = s.length();
        if (len <= 1) return s;
        return s = s.charAt(len - 1) + reverse(s.substring(1, len-1)) + s.charAt(0);
    }
    

    // Pre-condizioni: nessuna
    // Post-condizioni: restituisce la rappresentazione (String) di l
    // Effetti collaterali: nessuno   
    static String lToS(long l) {
        return "" + l;
    }
    
    
    // Pre-condizioni: s rappresenta un numero e non può essere null
    // Post-condizioni: restituisce la rappresentazione (long) di s
    // Effetti collaterali: nessuno
    static long sToL(String s) {
        return Long.parseLong(s);
    }

    
    // Pre-condizioni: s not null
    // Post-condizioni: true se s è palindroma.
    // Effetti collaterali: nessuno
    static boolean isPalin(String s) {
        int len = s.length();
        if (len <= 1) return true;
        return s.charAt(len - 1) == s.charAt(0) && isPalin(s.substring(1, len - 1));
    }

    
    // Pre-condizioni: n non è un numero di lychrel
    // Post-condizione: stampa la sequenza di lychrel
    // Effetti collaterali: System.out è modificato
    static void printLychrelSequence(long n) {
        while(!isPalin(lToS(n))) {
            System.out.println(n);
            n = nextLychrel(n);
        }
        System.out.println(n);
    }
    

    // Pre-condizioni: n >= 0
    // Post-condizioni: calcola il numero successivo della sequenza di lychrel
    // Effetti collaterali: nessuno
    static long nextLychrel(long n) {
        return n + sToL(reverse(lToS(n)));
    }

    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        long l = s.nextLong();
        printLychrelSequence(l);
        // Controllare se il numero n è palindromo
        // se lo è ci fermiamo
        // se non lo è lo robaltiamo
        // poi sommiamo i due numeri
        // e ripetiamo
    }
}
