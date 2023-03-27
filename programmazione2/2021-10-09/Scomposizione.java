import java.util.List;
import java.util.ArrayList;

public class Scomposizione {
    public static void main(String[] args) {
        int limit = Integer.parseInt(args[0]);
        int tmpNumber;
        List<Integer> scomp = new ArrayList<>();

        for (int number = 2; number <= limit; number++) {
            tmpNumber = number;
            for (int div = 2; div <= tmpNumber;) {
                if ((tmpNumber % div) == 0) {
                    scomp.add(div);
                    tmpNumber /= div;
                } else {
                    div++;
                }
            }
            System.out.println("\t" + number + " => "  + scomp);
            scomp.clear();
        }
    }
}
