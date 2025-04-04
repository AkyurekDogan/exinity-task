
select * from partner where id='pl12lll123'

select
	p.id,
	p.name,
	ST_X(p."location"::geometry) as lat,
	ST_Y(p."location"::geometry) as long,
	p.radius,
	ST_Distance(
        p."location",
        ST_SetSRID(ST_MakePoint(52.451320, 13.337652), 4326)::GEOGRAPHY
    ) AS distance,
    ps.craftsmanship_tags,
    pr.avg
from public.partner p
inner join public.skill ps on ps.partner_id = p.id
left join public.rating pr on pr.partner_id = p.id 
where ST_DWithin(
    p."location",            
    ST_SetSRID(ST_MakePoint(52.451320, 13.337652), 4326)::GEOGRAPHY,
    p.radius*1000
) and ps.craftsmanship_tags @> '{wood}'
order by pr.avg desc, distance asc

