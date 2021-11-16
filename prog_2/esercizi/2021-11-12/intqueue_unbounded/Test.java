import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;

public class Test {
    public static void main(String[] args) {
        int plus = 0;
        int minus = 0;
        int occ = 0;
        List<Integer> inputs = new ArrayList<>();
        Scanner s = new Scanner(System.in);
        UnboundedIntQueue queue = new UnboundedIntQueue();
        
        while(s.hasNextInt()) {
            inputs.add(s.nextInt());
        }

        for(int element : inputs) {
            if(element == 1) {
                plus++;
            } else if(element == -1) {
                minus++;
            }

            if(element == 1) {
                occ++;
                queue.enqueue(occ);
            } else if(element == -1) {
                System.out.println(queue.firstElement());
                queue.dequeue();
            }
            if (minus > plus) {
                System.out.println(queue.toString());
            }       
        }
        System.out.println(queue.toString());
        System.out.println(queue.dim());
    }
}
