import java.util.NoSuchElementException;

// INTERFACCIA
// Interfaccia utile per scegliere dei candidati per delle interrogazioni
public interface Picker {
	
	/*	Restituisce uno dei candidati possibili
	 *	
	 *	@return un candidato
	 *	@throws NoSuchElementException se non ci sono piu candidati
	 */
	public String pick() throws NoSuchElementException;

	/*	Restituisce un numero richiesto di candidati
	 *  
	 *	Se non ci sono abbastanza candidati soleva una {@code IllegalArgumentException}.
	 *
	 *	@param k il numero di candidati da restituire
	 *	@return array di k candidati
	 *	@throws IllegalArgumentException se k e' minore di 1
	 *	@throws NoSuchElementException se k eccede il numero di candidati rimanenti
	 */
	public default String[] pick(int k) throws IllegalArgumentException, NoSuchElementException {
		if(k < 1) throw new IllegalArgumentException("Il numero di candidati dev'essere positivo.");
		if(k > remaining()) throw new NoSuchElementException("Non ci sono abbastanza candidati.");
		final String[] retval = new String[k];
		for(int i = 0; i < k; i++) {
			retval[i] = pick();
		}
		return retval;
	}

	/*	Restituisce il numero dei candidati che sono rimasti da chiamare
	 *	
	 *	@return Numero dei candidati rimasti
	 */
	public int remaining();
}
