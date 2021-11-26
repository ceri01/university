import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;

public class Test {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        List<Rational> list = new ArrayList<Rational>();
        Rational r, sum, diff, mul;

        while(s.hasNext()) {
            r = new Rational(s.nextInt(), s.nextInt());
            list.add(r);
        }


        for(Rational rat : list) {
            r = rat.next();
            sum = rat.sum(r);
            diff = rat.diff(r);
            mul = rat.mul(r);
        }

        System.out.println("Sum: " + sum + "\ndiff: " + diff + "\nmul: " + mul + "\n");
    }
}
