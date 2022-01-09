import java.util.Collection;
import java.util.List;
import java.util.ArrayList;
import java.util.Comparator;

public class Squad {
	public final String name;
	private List<Solider> members;

	public Squad(List<Solider> s, String name) {
		if(s == null) throw new IllegalArgumentException("Lista di soldati nulla.");
		if(name == null) throw new IllegalArgumentException("Nome nullo.");
		for(Solider sol : s) {
			if(sol == null) throw new IllegalArgumentException("Elemento nella lista nullo");
		}
		this.name = name;
		this.members = new ArrayList() 
	}

	public void sort() {
		this.members.sort(new Comparator<>() {
			@Override
			public int compare(Jack s1, Jack s2) {
				return Integer.compare(s1.life(), s2.life());
			}
		});
	}

	public void show() {
		for(Solider s : this.members) {
			System.out.println("\t" + s.name);
		}
	}
}
