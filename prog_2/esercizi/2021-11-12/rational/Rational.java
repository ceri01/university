/*  Overview: Questa classe rappresenta un numero razionale, formato da due interi, un numeratore e
 *            un denominatore.
 *
 *            Un numero razionale e' immutabile dal momento che non dobbiamo modificatre direttamente
 *            il numero, ma se per esempio dobbiamo fare una somma tra due numeri razionali creeremo
 *            un oggetto che rappresenta il numero razionale risultante e lo restituiremo.
 */



public class Rational {
    /*
     *  La struttura dati usata per creare il numero razionale e' un record formato
     *  da due interi che rappresentano il numeratore e il denominatore
    */
    public record Rat(int num, int den) {
        // VA BENE METTERE IL RECORD A PUBLICI??
        public Rat{
            if(den == 0) throw new ArithmeticException("Non e' possibile avere il denominatore = 0");
        }
        
        @Override
        public String toString() {
            return num + "\n--\n" + den + "\n";
        }
    }

    /*
     *  ABS FUNC(rational): rational = +-(num)
     *                                 _______
     *                                 
     *                                 +-(den) 
     *
     *
     *  RAP INV: rational.den != 0, rational != NULL
     */

    // Attributi
    // VA BENE METTERE GLI ATTRIBUTI PUBLIC FINAL? SI PER OGGETTI IMMUTABILI MA ESSENDO IL RECORD PUBLIC CI SONO PROBLEMI?
    
    public final Rat rational;

    // Costruttori
    
    /*  
     *  Costruttore usato per creare il numero razionale
     *
     *  Pre-condizioni: Il denominatore deve essere diverso da zero
     *  Post-condizioni: Viene creato il numero razionale formatod da numeratore e denominatore
     */


    public Rational(int num, int den) {
        int mcd = mcd(num, den);
        rational = new Rat(num/mcd, den/mcd);
    }

    // Metodi

    /*  
     *  Metodo che permette di trovare l'mcd tra due numeri
     *
     *  Post-condizioni: Viene restituito l'mcd tra num e den
     */

    private int mcd(int num, int den) {
        int a = num;
        int b = den;
        int r;

        if(a < b) {
            int tmp = a;
            a = b;
            b = tmp;
        }
        
        while(b != 0) { 
            r = a % b;
            a = b;
            b = r;
        }

        return a;
    }

    /*  
     *  Metodo che permette di calcolare l'mcm tra deu numeri
     *
     *  Post-condizioni: Viene restituito l'mcm tra num e den
     */

    private int mcm(int num, int den) {
        return (num * den) / mcd(num, den);
    }
    
    
    /*  
     *  Metodo che permette di semplificare il numero razionale
     *  Pre-condizioni: den non deve essere 0
     *  Post-condizioni: Viene restituito un nuovo numero razionale che rappresenta il numero this semplificato
     */

    public Rational simplify() {
        int mcd = mcd(rational.num, rational.den);
        Rational new_rational = new Rational(rational.num/mcd, rational.den/mcd);
        return new_rational;
    }

    /*
     *  Metodo che permette di eseguira la somma tra due numeri razionali
     *
     *  Post-condizioni: Restituisce un nuovo numero razionale che rappresenta la somma tra due numeri razionali
     */

    public Rational sum(Rational n) {
        // SEMPRE PER LA STORIA DEL RECORD PUBLIC POSSO FARE n.rationali, VUOL DIRE ESPORRE LA RAPPRESENTAZIONE? CI POSSO ACCEDERE ANCHE DA FUORI LA CLASSE MA NON POSSO MODIFICARLO, QUINDI VA BENE LO STESSO O ESPONGO LA RAPPRESENTAZIONE?
        int mcm = mcm(rational.den, n.rational.den);
        int num = ((mcm / rational.den) * rational.num) + ((mcm / n.rational.den) * n.rational.num);
       
        Rational sum = new Rational(num, mcm);
        return sum;
    }

    /*
     *  Metodo che permette di eseguira la differenza tra due numeri razionali
     *
     *  Post-condizioni: Restituisce un nuovo numero razionale che rappresenta la differenza tra due numeri interi
     */

    public Rational diff(Rational n) {
        int mcm = mcm(rational.den, n.rational.den);
        int num = ((mcm / rational.den) * rational.num) - ((mcm / n.rational.den) * n.rational.num);
       
        Rational diff = new Rational(num, mcm);
        return diff;   
    }

    /*
     *  Metodo che permette di generare il reciproco di un numero razionale
     *
     *  Post-condizioni: Restituisce il reciproco di un numro razionale
     */

    public Rational reciprocal() {
        return new Rational(rational.den, rational.num);
    }

    /*
     *  Metodo che permette di eseguira la moltiplicazione tra due numeri razionali
     *
     *  Post-condizioni: Viene reatituito un nuovo numero razionale che rappresenta la moltiplicazione tra due numeri razionali
     */

    public Rational mul(Rational n) {
        return new Rational(rational.num * n.rational.num, rational.den * n.rational.den);
    }

    
    /*
     *  Metodo che permette di eseguira la divisione tra due numeri razionali
     *
     *  Post-condizioni: Viene reatituito un nuovo numero razionale che rappresenta la divisione tra due numeri razionali
     */

    public Rational div(Rational n) {
        return mul(n.reciprocal());
    }
    
    //  I METODI EQUALS ECC HANNO BISOGNI CHE VENGA SPECIFICATA PRE POST CONDIZIONE ECC

    /*
     *  Metodo che controlla se due numeri razionali sono uguali
     *
     *  Pre-condizioni: obj non deve essere NULL
     *  Post-condizioni: Viene restituito true se i due oggetti sono simili (non sono comunque lo stesso oggetto)
     */

    public boolean equals(Object obj) {
        if(!(obj instanceof Rational)) return false;
        Rational other = (Rational) obj;
        
        if(rational.num == obj.rational.num && rational.den == obj.rational.den) {
            return true;
        }
        return false;
    }

    /*  
     *  Metodo che genera l'hashcode di un numero razionale
     *
     *  Pre-condizioni:
     *  Effetti collaterali:
     *  Post-condizioni:
     */
    // @Override
    // public int hashCode() {}

    /*
     *  Metodo che permette di stampare un numero razionale
     *
     *  Effetti collaterali: Standard output viene modificato
     *  Post-condizioni: Ritorna il numero razionale sottoforma di stringa
     */

    public String toString() {
        return rational.toString();
    }

    /*
     *  Metodo che permette di verificare se l'invariante di rappresentazione e' rispettato
     *
     *  Post-condizioni: Ritorna true se l'invariante di rappresentazione viene rispettato, false altrimenti
     */

    private boolean repOk() {
        return rational.den != 0
            && rational != null;
    }

    /*
     *  Metodo che permette di verificare se l'invariante di rappresentazione e' rispettato
     *
     *  Post-condizioni: Ritorna true se l'invariante di rappresentazione viene rispettato, false altrimenti
     */

    public class Iterator {
        public hasNext() {}

        public Next() {}
    }

    public Iterator iterator() {
        return;
    }
}
