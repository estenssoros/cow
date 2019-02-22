import { createStore, applyMiddleware, compose } from "redux";
import thunk from 'redux-thunk'
import { persistStore, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'
import rootReducer from "../reducers/index";

const INITIAL_STATE = {}
const composeEnhancer = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
const enhancer = composeEnhancer(applyMiddleware(thunk))
const persistConfig = {
    key: 'root',
    storage
}
const persistedReducer = persistReducer(persistConfig, rootReducer)
const store = createStore(persistedReducer, INITIAL_STATE, enhancer)
const persistor = persistStore(store)
export { store, persistor };