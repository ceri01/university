import java.util.Scanner;

public class Test {
	public static void main(String[] args) {
		SimpleMap map = new SimpleMap();
		StringBuilder sb = new StringBuilder("");
		Scanner s = new Scanner(System.in);
		

		while(s.hasNextLine()) {
			sb.setLength(0); 
			sb = new StringBuilder(s.next());
			if((sb.toString().equals("+"))) {
				sb = new StringBuilder(s.next());
				int val = s.nextInt();
				map.put(sb.toString(), val);
				System.out.println(map.toString());			
			} else if(sb.toString().equals("-")) {
				String key = s.next();
				if(map.contains(key)) {
					int val = map.get(key);
					map.remove(key);
					System.out.println(val);
				}
			} else if (sb.toString().equals("f")) {
				break;
			}
		}
		System.out.println(map.size());
	}
}
