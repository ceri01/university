import java.util.Set;
import java.util.HashSet;
import java.util.Scanner;

public class Main {
	public static void main(String[] args) {	
		Scanner scanner = new Scanner(System.in);
		Set<Word> s = new HashSet<>();
		Anagrams m = new Anagrams();
		
		while(scanner.hasNext()) {
			Word w = new Word(scanner.nextLine());
			s.add(w);
		}

		m.insert(s);
		m.print();
	}
}
