import java.util.List;
import java.util.NoSuchElementException;
import java.util.Random;

// Una implementazione di {@code AbstractListPicker} che sceglie i candidati in sequenza
public class RandomListPicker extends AbstractListPicker{
    
	private final Random rnd = new Random();
	
    protected RandomListPicker(final List<String> candidates) {
        super(candidates);
    }   

    @Override
    public String pick() throws NoSuchElementException {
        if(remaining() == 0) throw new NoSuchElementException("Non sono rimasti candidati.");
        return remove(rnd.nextInt(remaining()));
    }   
}

