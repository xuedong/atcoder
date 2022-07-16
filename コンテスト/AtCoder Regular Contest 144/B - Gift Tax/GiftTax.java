import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.*;

public class Main {

	static class DoubleEndedPriorityQueue {

	    Set<Integer> set;

	    DoubleEndedPriorityQueue() {
	        set = new HashSet<Integer>();
	    }

	    int size() {
	        return set.size();
	    }

	    boolean isEmpty() {
	        return set.size() == 0;
	    }

	    void insert(int x) {
	        set.add(x);
	    }

	    int getMin() {
	        return Collections.min(set, null);
	    }

	    int getMax() {
	        return Collections.max(set, null);
	    }

	    void deleteMin() {
	        if (set.size() == 0)
	            return;
	        set.remove(Collections.min(set, null));
	    }

	    void deleteMax() {
	        if (set.size() == 0)
	            return;
	        set.remove(Collections.max(set, null));
	    }

	};

	public static void main(String[] args) {
		String str1, str2;
		try {
			BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
			str1 = br.readLine();
			str2 = br.readLine();
		} catch (Exception e) {
			System.out.println("Input Error!!");
			return;
		}

		String[] sa1 = str1.split(" ");
		int n = Integer.parseInt(sa1[0]);
		int a = Integer.parseInt(sa1[1]);
		int b = Integer.parseInt(sa1[2]);

		String[] sa2 = str2.split(" ");
		int[] nums = new int[n];
		for (int i = 0; i < n; i++) {
			String s = sa2[i];
			nums[i] = Integer.parseInt(s);
		}

		giftTax(nums, n, a, b);
	}

	private static void giftTax(int[] nums, int n, int a, int b) {
		DoubleEndedPriorityQueue pq = new DoubleEndedPriorityQueue();

		for (int i = 0; i < n; i++) {
			int curr = nums[i];
			pq.insert(curr);
		}

		int ans = Integer.MIN_VALUE;
		while (pq.getMin() > ans) {
			ans = pq.getMin();

			int min = pq.getMin();
			pq.deleteMin();
			if (pq.isEmpty()) {
				break;
			}

			int max = pq.getMax();
			pq.deleteMax();

			int div = 1;
			if (pq.isEmpty()) {
				int diff = max - min;
				div = diff / (a + b) + 1;
			} else {
				int nextMin = pq.getMin();
				int divMin = (nextMin - min) / a;
				if ((nextMin - min) % a > 0) {
					divMin++;
				}

				int nextMax = pq.getMax();
				int divMax = (max - nextMax) / b + 1;
				if ((max - nextMax) % b > 0) {
					divMax++;
				}

				div = Math.min(divMin, divMax);
			}
			max -= div * b;
			min += div * a;

			pq.insert(max);
			pq.insert(min);
		}

		System.out.println(ans);
	}
}
