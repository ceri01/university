package filesystem;

/*	Questa interfaccia rappresenta un elemento all'interno del filesystem. tramite Entry e' possibile implementare
 *	le classi File e Directory
 */

public interface Entry {
	
	/*	Metodo che permette di ricavare il nome di una entry
	 *	
	 *	Post-condizioni: viene restituita una stringa che rappresenta il nome della entry
	 */
	public String name();
		
	/*	Metodo che permette di ricavare la dimensione della entry
	 *	
	 *	Post-condizioni: viene restituito un intero che rappresenta la dimensione della entry
	 */

	public int size();

	/*	Metodo che permette di ricavare il path di una entry
	 *	
	 *	Post-condizioni: viene restituito un oggetto di tipo path che rappresenta il path della entry
	 */
	
	public Path path();

}
