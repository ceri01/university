import java.util.Set;
import java.util.TreeSet;
import java.util.Map;
import java.util.HashMap;
import java.util.Comparator;

/*	Questa classe rappresenta tutte le parole che hanno in comune le stesse lettere, ovvero degli anagrammi.
 *	
 *	Gli oggetti di questa classe sono mutabile in quanto e' possibile inserire delle nuove parole ogni volta.
 *	
 *	Per rappresentare tutti gli anagrammi la classe utilizza come chiave la firma delle parole e 
 *	come valore un insieme (di word) contenente tutte le parole aventi la stessa firma (ovvero gli anagrammi, 
 *	le parole che hanno in comune le lettere).
 *
 *	ABS FUN: AF(anagramsMap) = mappa avente come chiave la firma delle prole e come valore l'insieme degli
 *							   anagrammi (word).
 *							   {key_0 -> [anag0_0, anag0_1, ..., anag0_n],
 *							   	key_1 -> [anag1_0, anag1_1, ..., anag1_n],
 *							   	.....
 *							   	key_n -> [anagn_0, anagn_1, ..., anagn_n]}
 *	RAP INV: Quando viene chiamato il costruttore la mappa viene istanziata, ma non sono richiesti parametri
 *			 quindi non c'e' bisogno di alcun vincolo. l'unico vincolo che abbiamo e' sull'inserimento dei
 *			 dati, dove l'insieme non puo' essere nullo. 
 */

public class Anagrams {
	// Attributi
	private final Map<String, Set<Word>> anagramsMap; 
	
	// Costruttore
	public Anagrams() {
		anagramsMap = new HashMap<String, Set<Word>>();		
	}

	// Metodi
	
	/*	Metodo che permette la creazione di una mappa vuota
	 *
	 *	Post-condizioni: anagramsMap viene istanziata
	 *
	 */
	
	public void insert(Set<Word> ws) {
		if(ws == null) throw new  IllegalArgumentException("L'insieme delle parole non puo' essere nullo.");
		for(Word w : ws) {
			if(!anagramsMap.containsKey(w.sign())) {
				Set<Word> set = new TreeSet<Word>();
				set.add(w);
				anagramsMap.put(w.sign(), set);
			} else {
				anagramsMap.get(w.sign()).add(w);
			}
		}
	}

	/*	Metodo che permette la stampa della mappa
	 *	
	 *	Effetti collaterali: Standard output puo' cambiare
	 *	Post-condizioni: Stampa la mappa
	 */

	public void print() {
		for(String s : this.anagramsMap.keySet()) {
			Set<Word> vals = new TreeSet<>(anagramsMap.get(s));
			System.out.println("Chiave " + s + ":");
			System.out.println("\t" + vals);
		}
	}

}
