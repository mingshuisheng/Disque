import classnames, {Argument, Mapping} from "classnames"

const isString = (v: any): v is string => {
  return typeof v === "string"
}

const isArray = (v: any): v is Array<Argument> => {
  return v instanceof Array
}

const isMapping = (v: any): v is Mapping => {
  return v !== null && typeof v === "object"
}

export namespace ClassNameUtils {
  export const create = (prefix: string) => {
    return {
      rootClass: (propsClass?: string, ...rest: Array<Argument>): string => {
        return classnames(propsClass, addAllPrefix(prefix, "root",rest));
      },
      className: (...classNames: Argument[]) => {
        return addAllPrefix(prefix, classNames)
      }
    }
  }

  const addAllPrefix = (prefix: string, ...args: Array<Argument>): string => {
    return classnames(processArray(prefix, args))
  }

  const processArray = (prefix: string, args: Array<Argument>): Array<Argument> => {
    for (let i = 0; i < args.length; i++) {
      let arg = args[i];
      if(isString(arg)){
        args[i] = processString(prefix, arg)
      }else if(isArray(arg)){
        args[i] = processArray(prefix, arg)
      }else if(isMapping(arg)){
        args[i] = processMapping(prefix, arg)
      }
    }
    return args
  }

  const processString = (prefix: string, arg: string): string => {
    return `${prefix}-${arg}`
  }

  function processMapping(prefix: string, arg: Mapping): Mapping {
    let newArg:Mapping = {}

    Object.keys(arg).forEach(key => {
      let newKey = processString(prefix, key)
      newArg[newKey] = arg[key]
    })

    return newArg;
  }

}