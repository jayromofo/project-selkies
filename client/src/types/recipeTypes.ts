export type Recipe = {
   ID: number
   name: string;
   description: string;
   type: string;
   category: string;
   Instructions?: RecipeInstruction[];
   MetaData?: MetaData;
}


export type RecipeCategory = {
   id: number;
   name: string;
   description: string;
}

export type MetaData = {
   ID: number;
   recipe_id: number;
   servings: number;
   cook_time: number;
   is_keto: boolean;
   is_vegetarian: boolean;
   tags?: string;
   image_path?: string;
}

export type RecipeInstruction = {
   ID: number;
   recipe_id: number;
   Line_Num: number;
   instruction: string;
}

