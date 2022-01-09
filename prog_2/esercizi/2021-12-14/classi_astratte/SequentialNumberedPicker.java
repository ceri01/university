public class SequentialNumberedPicker implements Picker{
	
	private final int tot;
	private int next;

	/*	REP INV: 1) tot >= 0
	 *
	 *	ABS FUN: rappresentiamo i numeri da estrarre con una variabile tot che indica i numeri totali che
	 *			 abbiamo e una variabile next che indica il prossimo numero da estrarre
	 *
	 */

	public SequentialNumberedPicker(final int tot) {
		if(tot < 0) throw new IllegalArgumentException("Il numero di candidati non puo' essere minore di 0");
		this.tot = tot;
		next = 0;
	}
	
	@Override
	public String pick() throws IllegalArgumentException {
		final int retval = next;
		next += 1;
		return "" + retval;
	}

	@Override
	public int remaining() {
		return tot - next;
	}

	
}
