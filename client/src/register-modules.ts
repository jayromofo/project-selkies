import router from "./router"
// import store from "./store"

type Module = {
   router: any
   key: string
}

const registerModule = (name: string, module: Module) => {
   // if (module.store) {
   //    store.registerModule(name, module.store);
   // }

   if (module.router) {
      module.router(router)
   }
};

export const registerModules = (modules: any)  => {
   Object.keys(modules).forEach((moduleKey: string) => {
      const module: Module = modules[moduleKey];
      registerModule(moduleKey, module);
      console.log(`moduleKey: ${moduleKey}, module: ${module}`);
   });
};