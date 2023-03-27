/*  Overview: Questa classe rappresenta un numero razionale, formato da due interi, un numeratore e
 *            un denominatore.
 *
 *            Un numero razionale e' immutabile dal momento che non dobbiamo modificatre direttamente
 *            il numero, ma se per esempio dobbiamo fare una somma tra due numeri razionali creeremo
 *            un oggetto che rappresenta il numero razionale risultante e lo restituiremo.
 */

import java.lang.Math;
import java.lang.ArithmeticException;
import java.util.Objects;

public class Rational {
    /*  Attributi
     *  Per rappresentare il numero razionale utilizzo due interi.
    */
    

    private int num;
    private int den;

    /*
     *  ABS FUNC(rational): rational = +-(num)
     *                                 _______
     *                                 
     *                                 +-(den) 
     *
     *
     *  RAP INV: this.den != 0, gcd(this.num, this.den)
	 *  ABS INV: this.den != 0
     */

    // Costruttori
    
    /*  
     *  Costruttore usato per creare il numero razionale
     *
     *  Pre-condizioni: Il denominatore deve essere diverso da zero
     *  Post-condizioni: Viene creato il numero razionale formatod da numeratore e denominatore
     */


    public Rational(int num, int den) {
        if(den == 0) throw new ArithmeticException("Denominator must be > 0");	
		if(den < 0 && num < 0) {
			den = Math.abs(den);
			num = Math.abs(num);
		} else if(den < 0 && num > 0) {

			den = Math.abs(den);
			num *= -1;
		}
		int mcd = mcd(Math.abs(num), den);
        this.num = num / mcd;
        this.den = den / mcd;
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
        int mcd = mcd(this.num, this.den);
        return new Rational(this.num/mcd, this.den/mcd);
    }

    /*
     *  Metodo che permette di eseguira la somma tra due numeri razionali
     *
     *  Post-condizioni: Restituisce un nuovo numero razionale che rappresenta la somma tra due numeri razionali
     */

    public Rational sum(Rational n) {
        int mcm = mcm(this.den, n.den);
        int num = ((mcm / this.den) * this.num) + ((mcm / n.den) * n.num);
        return new Rational(num, mcm);
    }

    /*
     *  Metodo che permette di eseguira la differenza tra due numeri razionali
     *
     *  Post-condizioni: Restituisce un nuovo numero razionale che rappresenta la differenza tra due numeri interi
     */

    public Rational diff(Rational n) {
        int mcm = mcm(this.den, n.den);
        int num = ((mcm / this.den) * this.num) - ((mcm / n.den) * n.num); 

    	return new Rational(num, mcm);
    }

    /*
     *  Metodo che permette di generare il reciproco di un numero razionale
     *
     *  Post-condizioni: Restituisce il reciproco di un numro razionale
     */

    public Rational reciprocal() {
		try {
			return new Rational(this.den, this.num);
		} catch(Exception e) {
			System.out.println("Reciprocal unavailable.\n");
		}
		return this;
    }

    /*
     *  Metodo che permette di eseguira la moltiplicazione tra due numeri razionali
     *
     *  Post-condizioni: Viene reatituito un nuovo numero razionale che rappresenta la moltiplicazione tra due numeri razionali
     */

    public Rational mul(Rational n) {
        return new Rational(this.num * n.num, this.den * n.den);
    }

    
    /*
     *  Metodo che permette di eseguira la divisione tra due numeri razionali
     *
     *  Post-condizioni: Viene reatituito un nuovo numero razionale che rappresenta la divisione tra due numeri razionali
     */

    public Rational div(Rational n) {
        return mul(n.reciprocal());
    }
    
	@Override
    public boolean equals(Object obj) {
        if(!(obj instanceof Rational)) return false;
        Rational other = (Rational) obj;
        
        if(this.num == other.num && this.den == other.den) {
            return true;
        }
        return false;
    }
 
    @Override
    public int hashCode() {
		return Objects.hash(this.num, this.den);
	}

	@Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
		sb.append(this.num);
		if(this.den != 1) {
			sb.append('/').append(this.den);
		}
		return sb.toString();
    }

	private boolean repOk() {
        return this.den != 0
				&& mcd(this.num, this.den) == 1;
    } 
}
