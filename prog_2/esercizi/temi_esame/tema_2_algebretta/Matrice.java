package tema_2_algebretta;

/*	Questa interfaccia permett di implementare delle classi che rappresentano una matrice.
 */

public interface Matrice {
	// Metodi
		
	/*	Metodo che permette di restituire la dimensione della matrice.
	 *	
	 *	post-condizioni: restituisce la dimensione della matrice.
	 */

	public int dim();

	/*	Metodo che permette di prendere il valore della matrice in posizione x, y.
	 *
	 *	pre-consizioni: x >= 0 && x < dim
	 *					y >= 0 && y < dim
	 *	post-consizioni: ritorna il valore in posizione x, y
	 *
	 */

	public int val(final int x, final int y);

	/*	Metodo che permette di fare la somma matriciale.
	 *
	 *	pre-condizioni: m != null
	 *					m non contiene elementi nulli.
	 *	post-condizioni: restituisce una matrice che rappresenta la somma tra m e
	 *					 la matrice che usa il metodo.
	 */

	public Matrice piu(final Matrice m);

	/*	Metodo che permette di eseguire il prodotto per uno scalare.
	 *	
	 *	post-condizioni: restituisce una matrice che rappresenta il prodotto tra i e
	 *					 la matrice che usa il metodo.
	 *
	 */

	public Matrice per_scalare(final int i);

	/*	Metodo che permette di eseguire il prodotto tra due matrici.
	 *
	 *	pre-condizioni: m != null
	 *					m non contiene elemeti nulli
	 *	post-condizioni: restituisce una matrice che rappresenta il prodotto tra la matrice che
	 *					 usa il metodo per m.
	 */

	public Matrice per_matrice(final Matrice m);

	/*	Metodo che permette di eseguire il prodotto tra una matrice e un vettore.
	 *
	 *	pre-condizioni: v != null
	 *					la dimensione del vettore deve essere la stessa di quella matrice. 
	 *	post-condizioni: restituisce un vettore che rappresenta il prodotto tra la matrice e il vettore v.
	 */

	public class Vettore per_vettore(final Vettore v);
}
