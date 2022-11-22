export const debounce = (func: Function, wait: number) => {
  let timerId: ReturnType<typeof setTimeout>;
  return (...args: any[]) => {
    if (timerId) clearTimeout(timerId);
    timerId = setTimeout(() => {
      func(...args);
    }, wait);
  };
};
