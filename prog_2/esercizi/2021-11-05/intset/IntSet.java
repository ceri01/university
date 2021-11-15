import java.util.*;

/**
 *  Gli oggetti di questa classe rappresentano l'insieme dei numeri iteri
 *  Gli oggetti di questa classe sono mutabili
 *  All'interno dell'insieme non ci sono elementi ripetuti
 *  I valori vengono inseriti in modo ordinato.
 */
@SuppressWarnings("unchecked")
public class IntSet {
    
    // ATTRIBUTI
    
    /**
     *  L'attributo set e' una lista che rappresenta un insieme di interi.
     */
    
    private List<Integer> set;
    // COSTRUTTORE

    /**
     *  Costruttore per l'insieme vuoto
     */

    public IntSet() {
        set = new ArrayList();
    }

    // METODI
    
    /**
     *  Post-condizioni: Restituisce un oggetto contenente gli stessi elementi di set
     *  Effetti collaterali: viene creato un nuovo oggetto uguale a set
     */
    public List<Integer> Set() {
        List<Integer> cloneSet = new ArrayList();
        cloneSet.addAll(0, set);
        return cloneSet;
    }

    /**
     *  Questa funzione ordina in maniera crescente i valori quando vengono inseriti
     *  Pre-condizioni: len > 0
     *  Post-condizioni: viene inserito un valore in modo ordinato nell'insieme
     */

    private void insertionSort(int len, int n) {
        int val, find;
        set.add(n);

        for(int i = 1; i < len; i++) {
            val = set.get(i);
            find = i - 1;
            while(find >= 0 && set.get(i) > val) {
                set.add(find+1, find);
                find = find - 1;    
            }
            set.add(find+1, val);
        }
    }

    /**
    *  Questa funzione permette di inserire un elemenro nell'insieme
    *
    *  Pre-condizioni: n deve essere un numero intero
    *  Post-condizioni: n verra' inserito nell'insieme 
    */

    public void insert(int n) {
        if(dim() == 1) {
            set.add(n);
        } else {
            insertionSort(dim()+1, n);
        }    
    }

    /**
     *  Questa funzione rimuove l'elemento specificato dall'insieme
     *  
     *  Post-condizioni: se e' presente nella lista il valore viene rimosso
     */

    public void delete(int n) {
        set.remove(Integer.valueOf(n));
    }

    /**
     *  Questa funzione ritorna un elemento casuale dell'insieme
     *  
     *  Effetti collaterali: Creazione dell'oggetto rand
     *  Post-condizioni: viene restituito un valore casuale dell'insieme
     */

    public int choose() {
        if(set.isEmpty()) throw new EmptyException("L'insieme e' vuoto");
        Random rand = new Random();
        return set.get(rand.nextInt(dim()));
    }
    
    /**
     *  Questa funzione restituisce la dimensione dell'insieme
     *
     *  Post-condizioni: viene restituira la dimensione dell'insieme
     */

    public int dim() {
        return set.size();
    }

    /**
     *  Questa funzione ricerca un elemento specificato
     *
     *  Post-condizioni: ritorna la posizione dell'elemento specificato se presente, -1 altrimenti.
     */
    
    private int binarySearch(int x) {
        int min = 0;
        int max = dim();
        int middle;

        while(min <= max) {
            middle = (max + min) >>> 1;
            if(set.get(middle) == x) {
                return middle;
            }
            if (set.get(middle) > x) {
                min = middle + 1;
            } else {
                max = middle - 1;
            }
        }
        return -1;
    }

    /**
     *  Questa funzioene ritorna controlla se l'elemento specificato e' nell'insieme oppure no
     *  Post-condizioni: la funzione ritorna true se x e' presente nell'insieme, false altrimenti
     */

    public boolean contains(int x) {
        if (binarySearch(x) >= 0) {
            return true;
        }
        return false;
    }
    
    /**
     *  Questa funzione stampa il contenuto dell'insieme
     *
     *  Effetti collaterali: lo standard output viene modificato
     *  Post-condizioni: viene stampato il comtenuto dell'insieme
     */

    @Override
    public String toString() {
        return "{" + set.toString() + "}";
    }
}
