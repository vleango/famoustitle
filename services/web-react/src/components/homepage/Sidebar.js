import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Form, Input } from 'reactstrap';
import { size, reverse, map } from 'lodash';
import moment from 'moment';
import Spinner from '../shared/Spinner';

import fontawesome from '@fortawesome/fontawesome';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faSearch from '@fortawesome/fontawesome-free-solid/faSearch';

import './css/sidebar.css';

fontawesome.library.add(faSearch);

export class Sidebar extends Component {

    constructor(props) {
        super(props);
        this.state = {
            search: ""
        };
    }

    componentWillReceiveProps = (nextProps) => {
        if(nextProps.selected) {
            let search = "";
            if(nextProps.selected.match) {
                search = nextProps.selected.match
            }
            this.setState({
                search: search
            });
        }
    };

    onInputChange = (e) => {
        const field = e.target.name;
        const value = e.target.value;
        this.setState(() => ({ [field]: value }));
    };

    onSearchSubmit = (e) => {
        e && e.preventDefault();
        this.props.updateFilter("match", this.state.search);
    };

    render() {
        return (
            <div className="search-container">
                <aside>
                    <Form className="search-form" onSubmit={this.onSearchSubmit}>
                        <Input className="form-control pr-5" type="search" name="search" placeholder="Search..." value={this.state.search} onChange={this.onInputChange} />
                        <button className="search-button" type="submit"><FontAwesomeIcon icon="search"/></button>
                    </Form>
                </aside>
                <aside className="widget">
                    <div className="widget-title">Archives</div>
                    { size(this.props.archives) === 0 && <Spinner/> }
                    <ul>
                        {
                            this.props.archives && reverse(map(this.props.archives, (count, date) => {
                                const momentDate = moment.utc(date);
                                let className = "btn btn-link widget--date";
                                if(this.props.selected && this.props.selected.date === momentDate.format("YYYY-MM-DD") ? "* " : "") {
                                    className += " widget__selected";
                                }
                                return <li key={date}>
                                    <button className={className} onClick={() => this.props.updateFilter("date", momentDate.format("YYYY-MM-DD"))}>{momentDate.format("MMMM YYYY")} ({count})</button>
                                </li>
                            }))
                        }
                    </ul>
                </aside>
                <aside className="widget">
                    <div className="widget-title">Tags</div>
                    { size(this.props.archives) === 0 && <Spinner/> }
                    <div className="tagcloud">
                        {
                            this.props.tags && this.props.tags.map((tag) => {
                                return <button key={tag} className={this.props.selected && this.props.selected.tag === tag ? "selected" : ""} onClick={() => this.props.updateFilter("tag", tag)}>{tag}</button>
                            })
                        }
                    </div>
                </aside>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        archives: state.articles.index.archives,
        tags: state.articles.index.tags,
        selected: state.articles.index.selected
    };
};

export default connect(mapStateToProps)(Sidebar);
