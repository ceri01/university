import java.util.ArrayList;
import java.util.List;

public class Main {
	public static void main(String[] args) {
		Solider a = new Jack(30, "Luciano");
		Solider b = new Jack(12, "Grognar");
		Solider c = new Jack(32, "Wraps");
		Solider d = new Jack(3, "Gregorio");
		Solider e = new Jack(76, "Ciro");
		
		List sl = new ArrayList<Solider>();
		sl.add(a);
		sl.add(b);
		sl.add(c);
		sl.add(d);
		sl.add(e);
		
		Squad sq = new Squad(sl, "Gli acquilotti");
		
		sq.show();
		sq.sort();
		sq.show();
	}
}
