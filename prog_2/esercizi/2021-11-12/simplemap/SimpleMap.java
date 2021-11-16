import java.util.List;
import java.util.ArrayList;
import java.util.Iterator;

/*  Overview: Questa classe implementa una mappa, ovvero una struttura dati in cui per ogni chiave e' associato
 *            un elemento, in piu' le chiavi sono uniche, quindi non e' possibile avere chiavi duplicate.
 *            La classe e' mutabile, siccome e' possibile eseguire operazioni di inserimento, rimozione.
 *            esempio mappa:
 *                          [k1 -> v1]
 *                          [k2 -> v2]
 *                          [k4 -> v3]
 *                          [... -> ...]
 *                          [kk -> vk]                         
 */

public class SimpleMap {
    //Attributi
    
    /*
     *  La struttura dati usata per implementare la mappa e' formata da due liste dove per lo stesso indice
     *  abbiamo nella prima lista la chiave e nella seconda lista il valore.
     */

    private List<Integer> values;
    private List<String> keys; 
    
    /*
     *  ABS FUN: AF(keys, values)
     *           = keys[i] e values[i] | 0 <= i <= (keys.dim() | values.keys())
     *           = keys[], values[] liste vuote
     *
     *                  
     *
     *  REP INV: I due array devono avere lo stesso numero di elementi, di conseguenza hanno la stessa dimensione
     *           keys[i] e' unico.
     *           keys != NULL
     *           values != NULL
     *           keys[i] != NULL
     */
    
    //Costruttori
     
    /*
	 *	Costruttore che permette di creare una mappa vuota
	 *
     *  Post-Condizioni: Vengono create le liste che contengono le chiavi e i valori
     */

    public SimpleMap() {
		keys = new ArrayList<>();
		values = new ArrayList<>();
		assert repOk();
	}
    
    //Metodi
    
    /*
     *  Questa metodo permette l'inserimento di un elemento nella mappa
     *
     *  Pre-condizioni: key non deve essere NULL
     *  Effetti collaterali: Keys e values potrebbero venire modificati, se la chiave e' nulla viene sollevata un eccezione NullKeyException
     *  Post-Condizioni: Il valore con la relativa chiave vengono inseriti nella mappa
     */

    public void put(String key, int value) {
		
		assert repOk();
			
		if(key == null) throw new NullKeyException("La chiave inserita ha valore nullo.");
		if(!keys.contains(key)) {
			assert repOk();
			keys.add(key);
			values.add(value);
		}
	}

    /*
     *  Questa metodo ritorna true se e' presente nella mappa la chiave passata
     *
     *  Pre-condizioni: Key non deve essere NULL
     *  Effetti collaterali: Se la chiave e' nulla viene sollevata un eccezione NullKeyException 
     *  Post-Condizioni: Ritorna true se la chiave e' presente nella mappa 
     */
    
    public boolean contains(String key) {
		if(key == null) throw new NullKeyException("La chiave inserita ha valore nullo.");
		assert repOk();
		return keys.contains(key);
	}
    /*
     *  Questa metodo permette di rimuovere un elemento dalla mappa fornendo la chiave
     *
     *  Pre-condizioni: Key non deve essere NUll
     *  Effetti collaterali: Se la chiave e' nulla viene sollevata un eccezione NullKeyException    
     *  Post-Condizioni:
     */
    
    public void remove(String key) {
		if(key == null) throw new NullKeyException("La chiave inserita ha valore nullo.");
		if(keys.contains(key)) {
			assert repOk();
			int index = keys.indexOf(key);
			keys.remove(key);
			values.remove(index);
		}
	}
    
    /*
     *  Questa metodo permette di reperire il valore della chiave data
     *
     *  Pre-condizioni: Key non deve essere NULL
     *  Effetti collaterali: Se la chiave e' nulla viene sollevata un eccezione NullKeyException o NonExistentKeyException
     *  Post-Condizioni: Ritorna il valore associato alla chiave
     */

    public int get(String key) {
		if(key == null) throw new NullKeyException("La chiave inserita ha valore nullo.");
		if(!keys.contains(key)) throw new NonExistentKeyException("La chiave inserita non e' presente nella mappa"); 
		assert repOk();
		int index = keys.indexOf(key);
		return values.get(index);

	}

    /*
     *	Questa metodo restituisce la dimensione della mappa  
     *
     *  Effetti collaterali: Se la dimensione dei due array e' diersa sollevva l'eccezione InsubstantialMapException 
     *  Post-Condizioni: Viene ritornata la dimensione della lista keys (che e' uguale a quella della lista values)
     */
	
	public int size() {
		if(!(keys.size() == values.size())) throw new InsubstantialMapException("La mappa e' inconsistente");
		assert repOk();
		return keys.size();
	}

	/*
     *  Questa metodo verifica che venga rispettato l'invariante di rappresentazione
     *
     *  Effetti collaterali: moooolto inefficiente, non saprei come altro controllare
     *  Post-Condizioni: Ritorna true se l'invariante di rappresentazione viene rispettato
     */
	
    private boolean repOk() {
		for(String element : keys) {
			if(element == null) {
				return false;
			}
		}

		Iterator<String> iterKeys1 = keys.iterator();
		Iterator<String> iterKeys2 = keys.iterator();

		int c1 = 0;
		int c2 = 0;

		while(iterKeys1.hasNext()) {
			while(iterKeys2.hasNext()) {
				if(c1 == c2) {
					iterKeys2.next();
				} else {
					if (iterKeys1.next() == iterKeys2.next()) {
						return false;
					}
				}
				c2++;
			} 
			c1++;
		}
			
		return keys != null
			&& values != null;
	}

    /*
     *  Questa metodo permette di stampare la mappa
     *
     *  Effetti collaterali: Standard output viene modificato
     *  Post-Condizioni: La mappa viene stampata
     */

    @Override
    public String toString() {
		StringBuilder sb = new StringBuilder("Map:\n");
		if(!keys.isEmpty() && !values.isEmpty()) {
			Iterator<String> iterKeys = keys.iterator();
			Iterator<Integer> iterValues = values.iterator();
			while(iterKeys.hasNext() && iterValues.hasNext()) {
				sb.append("[" + iterKeys.next() + " -> ");
				sb.append(iterValues.next() + "]\n");
			}
		}
		return sb.toString();
	}

    /*
     *  Questo metodo permette di confrontare due mappe
     *
     *  Pre-condizioni: obj non deve essere null
     *  Post-Condizioni: Ritorna true se i due oggetti sono simili (hanno gli stessi elementi e la stessa lunghezza)
     */
    
    @Override
    public boolean equals(Object obj) {
		if(!(obj instanceof SimpleMap)) return false;
		SimpleMap other = (SimpleMap) obj;

		if(size() != other.size()) return false;
		
		assert repOk();

		Iterator<String> iterKeys1 = keys.iterator();
		Iterator<Integer> iterValues1 = values.iterator();
		
		Iterator<String> iterKeys2 = keys.iterator();
		Iterator<Integer> iterValues2 = values.iterator();

		while(iterKeys1.hasNext() && iterKeys2.hasNext() && iterValues1.hasNext() && iterValues2.hasNext()) {
			if(iterKeys1.next() == iterKeys2.next()) {
				return false;
			}
			if(iterValues1.next() == iterValues2.next()) {
				return false;
			}
		}

		return false;
	}

    /*
     * 	Questo metodo ritorna l'hash della mappa 
     *
     *  Post-Condizioni: Viene restituito il valore che corrisponde all'hash della mappa
     */
    
    public int hashcode() {
		return Integer.parseInt((keys.hashCode() + "" + values.hashCode()));
	}

}
