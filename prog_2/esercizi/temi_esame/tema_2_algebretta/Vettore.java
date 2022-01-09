package tema_2_algebretta

/*	Questa interfaccia permette di implementare delle classi che rappresentano un vettore
 */

public interface Vettore {
	// Metodi
	
	/*	Metodo che permette di recuperare la dimensione del vettore
	 *	
	 *	Post-condizioni: restituisce la dimensione del vettore
	 */
	
	public int dim();
	
	/*	Metodo che permette di recuperare l'elemento in i-esima posizione
	 *	
	 *	Pre-condizioni: i >= 0 && i < dim
	 *	Post-condizioni: restituisce l'elemento in i-esima posizione nel vettore
	 */
	
	public int val(final int i);

	/*	Metodo che permette di eseguire il prodotto del vettore per uno scalare alpha 
	 *	
	 *	Post-condizioni: restituisce il vettore contenente il prodotto tra alpha e il vettore
	 */

	public Vettore per(final int alpha);

	/*	Metodo che permette di eseguire la somma vettoriale tra il vettore e un altro
	 *	vettore v passato come argomento.
	 *	
	 *	Pre-condizioni: v != null
	 *					v non contiene elementi nulli
	 *	Post-condizioni: restituisce il vettore contenente la somma vettoriale dei due vettori
	 */

	public Vettore piu(final Vettore v);
}
