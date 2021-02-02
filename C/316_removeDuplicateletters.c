
void util(char *s, int *flag, char **result) {
  if (!*s)
    return;
  if (!flag[*s - 97]) {
    strcat(*result, *s);
  }
  util(s + 1, flag, result);
}

char *removeDuplicateLetters(char *s) {
  int flag[26] = {0};
  char *result = (char *)calloc(26, sizeof(char));
  util(s, flag, result);
  return result;
}
