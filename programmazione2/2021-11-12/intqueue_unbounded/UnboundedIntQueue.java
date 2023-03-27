import java.util.List;
import java.util.LinkedList;

/* OVERVIEW: Le istanze di questa classe rappresentano una coda illimitata di interi, gli elementi nella coda
 *           vengono inseriti mediante una politica FIFO (First In First Out).
 *           Gli oggetti di questa classe sono mutabili, siccome e' possibile aggiungere e rimuovere elementi
 *           dalla mia coda.
 *           esempio di una coda: [e1, e2, e3, e4, ..., ek].
 *           esempio di insereimento (queue): [e1, e2, e3, e4, ..., ek, ek+1].
 *           esempio di rimozione dalla coda (dequeue): [e2, e3, e4, ..., ek, ek+1].
 */

public class UnboundedIntQueue {
    // ATTRIBUTI
    
    // Come struttura dati per implementare la coda viene utilizzata una linked list 
    private List<Integer> queue;
    
    // Indici che rappresentano il primo e l'ultimo elemento della coda
    private int first;
    private int last;
    
    /*
     *  ABS FUN: AF(queue, first, last)
     *           = queue[i] con first <= i <= last
     *           = [queue[first], queue[first + 1], ..., queue[last-1], queue[last]] con -1 < first <= last
     *           = queue vuota se first = -1 e last = -1
     *
     *  REP INV: La coda contine un numero di elementi maggiore o uguale a 0
     *           !((first = -1 && last != -1) || (first != -1 || last = -1))
     *           first = -1 || first = 0
     *           last = -1 || last = dim() - 1
     *
     */

    // COSTRUTTORI
    
    /*  
     *  Costruttore della classe che genera una coda vuota
     *
     *  Post-condizione: Viene inizializzato queue in modo da rappresentare una una coda vuota, vengonon anche
     *                   settati i due indici first e last a -1
     */

    public UnboundedIntQueue() {
        queue = new LinkedList<Integer>();
        first = -1;
        last = -1;
    }

    // METODI
    
    /*
     *  Quesata funzione permette di conoscere la dimensione della coda.
     *
     *  Post-condizione: viene ritornato il valore della dimensione della coda.
     */

    public int dim() {
        if(first == -1 && last == -1) {
            return 0;
        } else {
            return last + 1;
        }
    }

    /*
     *  Quesata funzione permette di conoscere il valore che verra' estratto se venisse fatta una dequeue
     *
     *  Post-condizione: viene ritornato il primo valore della coda.
     */
    
    public int firstElement() {
        return queue.get(first);
    }

    /*
     *  Questa funzione permettere di conoscere se la coda e' vuota oppure no
     *  
     *  Post-condizione: Ritorna true se la coda e' vuota, false altrimenti
     */

    public boolean isEmpty() {
        return first == -1 && last == -1;
    }

    /* 
     *  Questa funzione permette di inserire un valore nella coda
     *
     *  Effetti collateral: Il valore degli indici viene modificato, la coda viene modificata
     *  Post-condizione: Viene inserito un elemento nella coda
     */
    
    public void enqueue(int n) {
        if(queue.isEmpty()) {
            first++;
        }
        queue.add(n);
        last++;        
    }
    
    /*
     *  Questa funzione premette di rimuovere l'elemento all'inizio della coda (FIFO)
     *
     *  Effetti collateral: Se la coda e' vuota viene sollevata un eccezione di tipo EmptyQueueException.
     *  Post-condizione: 
     */

    public void dequeue() {
        if(queue.isEmpty()) throw new EmptyQueueException("Impossibile rimuovere l'elemento, la coda e' vuota.");
        queue.remove(first);
        last--;
    }

    /*
     *  Post-condizione: Restituisce true se l'invariante di rappresentazione e' valida
     */

    private boolean repOk() {
        return !((first == -1 && last != -1) || (first != -1 || last == -1))
            && first == -1 || first == 0
            && last == -1 || last == dim() - 1
            && dim() >= 0;
    }

    /*
     *  Effetti collateral: Lo standard out puo' venire modificato
     *  Post-condizione: viene ritornata coda sottoforma di stringa
     */

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder("intQueueUnbounded: [");
        
        if(!isEmpty()) {
            int i = first;
            for(int el : queue) {
                sb.append(el);
                if(i == last) {
                    break;
                }
                i++;
                sb.append(", ");
            }
        }
        sb.append("]");
        return sb.toString();
    }

}
