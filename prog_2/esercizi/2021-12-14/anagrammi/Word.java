import java.util.Arrays;

/*	Questa classe rappresenta le parole in ingresso, ovvero degli oggetti non modificabili che hanno un metodo per fornire  
 *	firma della parola in ingresso (firma = stringa formata dalle lettere della parola in ordine alfabetico).
 *	Oltre a questo metodo la classe parola permette di conoscere la lunghezza della parola.
 *
 *	Una parola puo' essere formata solo da lettere e in piu' una parola non puo' essere null
 *
 *	ABS FUN: AF(word) = attributo String che rappresenta la parola
 *
 *	REP INV: word != null
 *			 la parola non contiene simboli o numeri
 *			 la parola deve avere almeno un carattere.
 */

public class Word {
	// Attributi
	private final String word;
	
	// Costruttori
	
	/*	Costruttore della classe Word.
	 *
	 *	Pre-condizioni: w non puo' contenere numeri e simboli, non puo' essere vuota e non puo' essere null
	 *					sollenva IllegalArgumentException se si verificano queste condizioni.
	 *	Post-condizioni: inizializza this.parola in modo che rappresenti la parola.
	 *
	 */
	public Word(final String w) {
		if(w == null) throw new IllegalArgumentException("La parola inserita non puo' essere null."); 
		if(w == "") throw new IllegalArgumentException("La parola non puo' essere vuota");
		for(char c : w.toCharArray()) {
			if(!Character.isLetter(c)) throw new IllegalArgumentException("La parola non puo' contenere numeri o simboli.");
		}
		this.word = w;
	}

	// Metodi
	
	/*	Metodo che ritorna la firma della parola.
	 *
	 *	Post-condizioni: Viene ritornata una stringa che rappresenta la firma della parola (formata dalle
	 *					 lettere della parola in ordine alfabetico).
	 *
	 */

	public String sign() {
		String sign = ""; // Stringa da restituire
		char[] tmp = this.word.toCharArray(); // Crea un array di char
		Arrays.sort(tmp); // Ordinamento dell'array di char
		sign = String.valueOf(tmp); // partendo dall'array ordinato creo la stringa da ritornare
		return sign;
	}

	/*	Metodo che restituisce la dimensione della parola.
	 *	
	 *	Post-condizioni: restituisce un intero che rappresenta la dimensione della parola.
	 */

	public int size() {
		return this.word.length();
	}

	@Override
	public String toString() {
		return this.word.toString();
	}

	@Override
	public int compare(Word w1, Word w2, new Comparator<>() {
		public new
	}) {
	
	}


}
