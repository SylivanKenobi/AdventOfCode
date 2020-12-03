import java.util.ArrayList;
import java.util.Arrays;

public class Advent02 {
    private static Integer[] opcode = {1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 19, 10, 23, 1, 23, 6, 27, 1, 6, 27, 31, 1, 13, 31, 35, 1, 13, 35, 39, 1, 39, 13, 43, 2, 43, 9, 47, 2, 6, 47, 51, 1, 51, 9, 55, 1, 55, 9, 59, 1, 59, 6, 63, 1, 9, 63, 67, 2, 67, 10, 71, 2, 71, 13, 75, 1, 10, 75, 79, 2, 10, 79, 83, 1, 83, 6, 87, 2, 87, 10, 91, 1, 91, 6, 95, 1, 95, 13, 99, 1, 99, 13, 103, 2, 103, 9, 107, 2, 107, 10, 111, 1, 5, 111, 115, 2, 115, 9, 119, 1, 5, 119, 123, 1, 123, 9, 127, 1, 127, 2, 131, 1, 5, 131, 0, 99, 2, 0, 14, 0};

    public static void main(String[] args) {
        for (int outer = 0; outer < 100; outer++) {
            for (int inner = 0; inner < 100; inner++) {
                ArrayList<Integer> opcodeList = new ArrayList(Arrays.asList(opcode));
                opcodeList.set(1, outer);
                opcodeList.set(2, inner);
                for (int i = 0; i < opcodeList.size(); i += 4) {
                    switch (opcodeList.get(i)) {
                        case 99:
                            i = opcodeList.size();
                            break;
                        case 1:
                            opcodeList.set(opcodeList.get(i + 3), opcodeList.get(opcodeList.get(i + 1)) + opcodeList.get(opcodeList.get(i + 2)));
                            break;
                        case 2:
                            opcodeList.set(opcodeList.get(i + 3), opcodeList.get(opcodeList.get(i + 1)) * opcodeList.get(opcodeList.get(i + 2)));
                            break;
                    }
                }
                if (opcodeList.get(0) == 19690720) {
                    System.out.println(opcodeList.get(1) + " " + opcodeList.get(2));
                    System.out.println(100 * outer + inner);
                }
            }
        }
    }
}

