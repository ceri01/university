import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;
import java.util.ListIterator;

public class Test {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        List<Rational> list = new ArrayList<Rational>();
        Rational r = null, sum = null, diff = null, mul = null, div = null;

        while(s.hasNext()) {
            r = new Rational(s.nextInt(), s.nextInt());
            list.add(r);
        }

		ListIterator<Rational> iter = list.listIterator();
		sum = iter.next();
		diff = sum;
		mul = sum;
        div = sum;

		while(iter.hasNext()) {	
			r = iter.next();
            sum = sum.sum(r);
            diff = diff.diff(r);
            mul = mul.mul(r);
			div = div.div(r);
        }
		
		System.out.println("Sum: " + sum.toString() + "\ndiff: " + diff.toString() + "\nmul: " + mul.toString() + "\ndiv: " + div.toString());
    }
}
