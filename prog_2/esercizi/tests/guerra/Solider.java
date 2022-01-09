import java.util.Map;

public interface Solider {
	
	/*	Method that allows the solider to attacks an enemy
	 *	
	 *	pre: s can't be null
	 *		 s can't be in the same squad of this
	 */
	
	public void attack(Solider s);
	
	/*	Method that allows the solider to get damages
	 *	
	 *	pre: damage >= 0
	 */

	public void getDamage(int damage);


	/*	Method that makes available solider's stats at the caller
	 *
	 *	post: returns a map whit: key = stat's name (string)
	 *							  value = value of stat
	 *
	 */

	public Map<String, Integer> getStats();

	public void kill();

	public int life();
}
