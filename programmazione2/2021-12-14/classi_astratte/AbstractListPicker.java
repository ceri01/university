import java.util.List;
import java.util.LinkedList;
import java.util.Objects;

// CLASSE ASTRATTA
// Implementazione parziale dell'interfaccia Picker che seleziona i candidati da una lista
public abstract class AbstractListPicker implements Picker {
	// Candidati da chiamare
	private final List<String> candidates;

	/*	REP INV: 1) candidates != NULL
	 *			 2) c !=  NULL per ogni c in candidates
	 *	
	 *	ABS FUN: candidati = [c0, c1, c2, ..., cn] rappresentiamo i candidati con una lista di candidati
	 */
	
	/*	Costruisce un istanza a partire da un elenco di candidati
	 *
	 *	@param candidates un elenco di candidati
	 *	@throws NUllPointerException se l'elenco e' null o contiene null
	 *	
	 *	protected perche' cosi possiamo utilizzarlo sono nelle sottoclassi e non farlo usare all'utente che utilizza il nostro programma
	 */
	protected AbstractListPicker(final List<String> candidates) throws NullPointerException {
		Objects.requireNonNull(candidates); // 1) rispettato
	   	for(final String c : candidates) {
			Objects.requireNonNull(c); // 2) rispettato
		}
		this.candidates = new LinkedList<>(candidates); // costruttore copia di LinkedList, cosi non passiamo direttamente il riferimento della lista in modo da preservare l'integrirta' dei dati
	}

	@Override
	public int remaining() {
		return candidates.size();
	}

	/*	Questo metodo restituisce l'i-esimo candidato e trasla di una posizione a sinistra i successivi
	 *
	 * 	@param i l'indice dell'elemento da rimuovere
	 * 	@throws IndexOutOfBoundsException se l'indice non corrisponde a nessun candidato rimanente
	 *	
	 *	protected cosi posso usarlo solo nella sottoclasse, il client non puo' togliere elementi.
	 *
	 *	final nei parametrio formali serve a chi scrive il metodo per ricordarsi che uso quella variabile
	 *	solo in lettura, e quindi non deve modificarla. e' un ESPRESSIONE DI UN INTENZIONE
	 */

	protected String remove(final int i) throws IndexOutOfBoundsException {
		return candidates.remove(i);
	}
}
