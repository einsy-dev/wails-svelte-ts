declare global {
  namespace svelteHTML {
    interface HTMLAttributes<T> {
      "on:click_outside"?: (event: CustomEvent) => void;
    }
  }
}

export {};
