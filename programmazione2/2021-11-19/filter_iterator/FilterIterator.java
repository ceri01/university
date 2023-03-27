import java.util.Iterator;
import java.util.function.Predicate;

public class FilterIterator implements Iterator {
    private final Iterator<T> iterator; // mettiamo final perche'
    private final Predicate<T> predicate;
    private boolean hasNext;
    private T next;

    FilterIterator(Iterator<T> iterator, final Predicate<T>, predicate) { // il final che c'e' qui vuol dire che non puo cambiare il riferimento a quell'oggetto, e si fa perche si evita di avere effetti collaterali a cio che si passa qui dentro. evitare di cambiare i riferimenti nei metodi di produzione
        this.iterator = iterator;
        this.predicate = predicate;
    }

    @Override
    public boolean hasNext() {
        while(!hasNext && iterator.hasNext()) {
            next = iterator.next();
            hasNext = predicate.test(next);
        }
        return hasNext;
    }

    @Override
    public T next() {
        if(!hasNext()) throw new NoSuchElementException("kekw");
        hasNext = false;
        return next;
    }
}
