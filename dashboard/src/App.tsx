import React from "react";
import { Link, Redirect, Route, Switch } from "react-router-dom";
import { ListContainer, ListImage, NewContainer } from "./views";

export default () => (
    <div>
        <nav className={'navbar navbar-expand-lg navbar-dark'}>
            <div className={'container'}>
                <Link to={'/'} className={'navbar-brand'}>Expected.sh</Link>
                <button className={'navbar-toggler'} type={'button'}>
                    <span className={'navbar-toggler-icon'} />
                </button>
                <div className={'collapse navbar-collapse'}>
                    <ul className={'navbar-nav'}>
                        <li className={'nav-item active'}>
                            <Link to={'/containers'}
                                className={'nav-link'}>Containers</Link>
                        </li>
                        <li className={'nav-item'}>
                            <Link to={'/images'} className={'nav-link'}>Images</Link>
                        </li>
                    </ul>
                    <ul className={'navbar-nav ml-auto'}>
                        <li className={'nav-item'}>
                            <Link to={'/settings'} className={'nav-link'}>Settings</Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>

        <div className={'container'}>
            <Switch>
                <Route path={'/containers/new'} component={NewContainer} />
                <Route path={'/containers'} component={ListContainer} />
                <Route path={'/images'} component={ListImage} />
                <Redirect from={'/'} to={'/containers'} />
            </Switch>
        </div>
    </div>
);
