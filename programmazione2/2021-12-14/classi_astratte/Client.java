import java.util.Arrays;
import java.util.List;
import java.util.Objects;

public class Client {
	public static void main(String[] args) {
		final List<String> candidates = List.of("Anna", "Paolo", "Mimmo", "Giovanni", "Ruben", "Aldo", "Ciro", "Mirko");
		Picker slp = new SequentialListPicker(candidates);
		Picker rlp = new RandomListPicker(candidates);
		Picker snp = new SequentialNumberedPicker(8);
		Picker rnp = new RandomNumberedPicker(8);

		// slp
		System.out.println("\n### slp ###");
		while(slp.remaining() >= 1) {
			System.out.println(slp.pick());
		}
		
		// rlp
		System.out.println("\n### rlp ###");
		while(rlp.remaining() >= 1) {
			System.out.println(rlp.pick());
		}
		
		// snp
		System.out.println("\n### snp ###");
		while(snp.remaining() >= 1) {
			System.out.println(snp.pick());
		}
		
		// rnp
		System.out.println("\n### rnp ###");
		while(rnp.remaining() >= 1) {
			System.out.println(rnp.pick());
		}
	}
}
