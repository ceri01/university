import java.util.Random;
import java.util.List;
import java.util.ArrayList;

public class RandomNumberedPicker implements Picker{
	
	private final int tot;
	private int next;
	private ArrayList<Boolean> picked; // lista dei numeri estratti
	private int totPicked; // totale numeri gia estratti
	private Random rnd = new Random();

	/*	REP INV: 1) tot >= 0
	 *
	 *	ABS FUN: rappresentiamo i numeri da estrarre con una variabile tot che indica i numeri totali che
	 *			 abbiamo e una variabile next che indica il prossimo numero da estrarre
	 */

	public RandomNumberedPicker(final int tot) {
		if(tot < 0) throw new IllegalArgumentException("Il numero di candidati non puo' essere minore di 0");
		picked = new ArrayList<Boolean>(tot);
		for(int i = 0; i < tot; i++) {
			picked.add(false);
		}
		this.tot = tot;
		next = rnd.nextInt(tot);
		totPicked = 0;
	}
	
	@Override
	public String pick() throws IllegalArgumentException {
		int retval = next;
		while(picked.get(next)) {
			next = rnd.nextInt(this.tot);
		}
		totPicked += 1;
		picked.set(next, true);
		retval = next;
		return "" + retval;
	}

	@Override
	public int remaining() {
		return tot - totPicked;
	}
}
