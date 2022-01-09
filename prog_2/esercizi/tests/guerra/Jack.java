import java.util.Map;
import java.util.HashMap;

public class Jack implements Solider {
	
	public static final Map<String, Integer> stats = new HashMap<String, Integer>() {{
		this.put("atk", 20);
		this.put("def", 35);
		this.put("str", 50);
		this.put("arm", 30);
	}};

	public final String name; // name of Jack	
	private int lp; // life points of Jack
	private int exp; // exp of Jack
	private Squad squad; // squad of Jack
	private boolean dead;
	
	public Jack(int exp, String name) {
		if(exp < 0) throw new IllegalArgumentException("exp can't be negative.");	
		if(name != null) throw new IllegalArgumentException("name is null");
		this.exp = exp;
		this.name = name;
		this.lp = 100;
		this.dead = false;
	}

	public void attack(Solider s) {
		if(s == null) throw new IllegalArgumentException("s can't be null.");
 		if(this.equals(s)) throw new IllegalArgumentException("impossible attaks own teammate.");
		s.getDamage(this.stats.get("atk"));
	}

	public void getDamage(int damage) {
		if(damage < 0) throw new IllegalArgumentException("damage can't be negative");
		this.lp = this.lp - damage;
		if(this.lp <= 0) {
			this.kill();
		}
	}

	public void kill() {
		this.dead = true;
	}

	public Map<String, Integer> getStats() {
		Map st = new HashMap(this.stats);
		return st;
	}

	public int life() {
		int l = this.lp;
		return l;
	}
}
