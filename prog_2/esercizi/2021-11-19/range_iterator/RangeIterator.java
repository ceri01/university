import java.util.Iterator;
import java.util.NoSuchElementException;

public class RangeIterator implements Iterable<Integer> {

  private final int to;
  private final int step;
  private int next;
  private boolean hasNext;

  RangeIterator(int from, int to, int step) {
    if (step == 0) throw new IllegalArgumentException("Step must be not 0");

    this.to = to;
    this.step = step;

    next = from;
    hasNext = step > 0 ? next < to : next > to;
  }

  RangeIterator(int to, int step) {
    this(0, to, step);
  }

  RangeIterator(int to) {
    this(to, 1);
  }

  @Override
  public Iterator<Integer> iterator() {
    return new Iterator<Integer>() {
      @Override
      public Integer next() {
        if (!hasNext()) throw new NoSuchElementException();
        hasNext = false;
        return next;
      }

      @Override
      public boolean hasNext() {
        if (!hasNext) {
          next += step;
          hasNext = step > 0 ? next < to : next > to;
        }
        return hasNext;
      }
    };
  }
}
