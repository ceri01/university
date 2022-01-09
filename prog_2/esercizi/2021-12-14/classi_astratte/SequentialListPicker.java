import java.util.List;
import java.util.NoSuchElementException;

// Una implementazione di {@code AbstractListPicker} che sceglie i candidati in sequenza
public class SequentialListPicker extends AbstractListPicker{
	
	protected SequentialListPicker(final List<String> candidates) {
		super(candidates);
	}

	@Override
	public String pick() throws NoSuchElementException {
		if(remaining() == 0) throw new NoSuchElementException("Non sono rimasti candidati.");
		return remove(0);
	}
}
