#include <cstdio>

int main () {
  int N, M, a, b, c[1000010];
  scanf("%d", &N);

  for (int k = 0; k < N; k++) {
    if (k > 0) printf("\n");

    scanf("%d", &M);
    for (int i = 0; i < M; i++) {
      scanf("%d %d", &a, &b);
      c[i] = a + b;
    }
    for (int i = M - 1; i > 0; i--) {
      if (c[i] >= 10) {
        c[i - 1] += 1;
        c[i] -= 10;
      }
    }
    for (int i = 0; i < M; i++) {
      printf("%d", c[i]);
    }
    printf("\n");
  }
}